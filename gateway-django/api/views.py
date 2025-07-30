import requests
from zeep import Client, Settings
from zeep.transports import Transport
from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework import status
from django.conf import settings
import logging

# Configurar logging
logger = logging.getLogger(__name__)

# URL base do backend Go
GO_BACKEND_URL = getattr(settings, 'GO_BACKEND_URL', 'http://localhost:8080')

@api_view(['GET', 'POST'])
def livros_view(request):
    """
    Gateway para endpoints de livros
    GET: Lista todos os livros
    POST: Cria um novo livro
    """
    try:
        if request.method == 'GET':
            # Fazer requisição para o backend Go
            response = requests.get(f'{GO_BACKEND_URL}/livros')
            
            if response.status_code == 200:
                return Response(response.json(), status=status.HTTP_200_OK)
            else:
                return Response(
                    {'error': 'Erro ao buscar livros no backend'}, 
                    status=response.status_code
                )
                
        elif request.method == 'POST':
            # Fazer requisição POST para o backend Go
            response = requests.post(
                f'{GO_BACKEND_URL}/livros/',
                json=request.data,
                headers={'Content-Type': 'application/json'}
            )
            
            if response.status_code in [200, 201]:
                try:
                    data = response.json()
                except ValueError:
                    data = request.data  
                return Response(data, status=status.HTTP_201_CREATED)
            else:
                return Response(
                    {'error': 'Erro ao criar livro no backend'}, 
                    status=response.status_code
                )
                
    except requests.exceptions.ConnectionError:
        logger.error("Erro de conexão com o backend Go")
        return Response(
            {'error': 'Não foi possível conectar ao backend. Verifique se o servidor Go está rodando na porta 8080.'}, 
            status=status.HTTP_503_SERVICE_UNAVAILABLE
        )
    except Exception as e:
        logger.error(f"Erro inesperado: {str(e)}")
        return Response(
            {'error': 'Erro interno do servidor'}, 
            status=status.HTTP_500_INTERNAL_SERVER_ERROR
        )

@api_view(['GET', 'POST'])
def usuarios_view(request):
    """
    Gateway para endpoints de usuários
    GET: Lista todos os usuários
    POST: Cria um novo usuário
    """
    try:
        if request.method == 'GET':
            response = requests.get(f'{GO_BACKEND_URL}/usuarios')
            
            if response.status_code == 200:
                return Response(response.json(), status=status.HTTP_200_OK)
            else:
                return Response(
                    {'error': 'Erro ao buscar usuários no backend'}, 
                    status=response.status_code
                )
                
        elif request.method == 'POST':
            response = requests.post(
                f'{GO_BACKEND_URL}/usuarios',
                json=request.data,
                headers={'Content-Type': 'application/json'}
            )
            
            if response.status_code in [200, 201]:
                try:
                    data = response.json()
                except ValueError:
                    data = request.data  # ou um dicionário com mensagem de sucesso
                return Response(data, status=status.HTTP_201_CREATED)
            else:
                return Response(
                    {'error': 'Erro ao criar usuário no backend'}, 
                    status=response.status_code
                )
                
    except requests.exceptions.ConnectionError:
        logger.error("Erro de conexão com o backend Go")
        return Response(
            {'error': 'Não foi possível conectar ao backend. Verifique se o servidor Go está rodando na porta 8080.'}, 
            status=status.HTTP_503_SERVICE_UNAVAILABLE
        )
    except Exception as e:
        logger.error(f"Erro inesperado: {str(e)}")
        return Response(
            {'error': 'Erro interno do servidor'}, 
            status=status.HTTP_500_INTERNAL_SERVER_ERROR
        )

@api_view(['GET', 'POST'])
def emprestimos_view(request):
    """
    Gateway para endpoints de empréstimos
    GET: Lista todos os empréstimos
    POST: Cria um novo empréstimo
    """
    try:
        if request.method == 'GET':
            response = requests.get(f'{GO_BACKEND_URL}/emprestimos')
            if response.status_code == 200:
                return Response(response.json(), status=status.HTTP_200_OK)
            else:
                return Response(
                    {'error': 'Erro ao buscar empréstimos no backend'}, 
                    status=response.status_code
                )
        elif request.method == 'POST':
            # Chamada SOAP para criar empréstimo
            try:
                wsdl_url = f'{GO_BACKEND_URL}/soap/biblioteca.wsdl'
                client = Client(wsdl=wsdl_url, settings=Settings(strict=False), transport=Transport(timeout=10))
                livro_id = request.data.get('livro_id')
                usuario_id = request.data.get('usuario_id')
                logger.error(f"DEBUG: usuario_id={usuario_id}, livro_id={livro_id}")
                # Chama o método SOAP (ajuste o nome conforme seu WSDL)
                response = client.service.EmprestarLivro(livro_id=livro_id, usuario_id=usuario_id)
                # Retorna a resposta SOAP como JSON para o frontend
                return Response({
                    'mensagem': getattr(response, 'Mensagem', ''),
                    'sucesso': getattr(response, 'Sucesso', False),
                    'emprestimo_id': getattr(response, 'EmprestimoID', None)
                }, status=status.HTTP_201_CREATED)
            except Exception as e:
                logger.error(f"Erro ao consumir SOAP: {str(e)}")
                return Response({'error': 'Erro ao comunicar com o backend SOAP'}, status=status.HTTP_500_INTERNAL_SERVER_ERROR)
    except requests.exceptions.ConnectionError:
        logger.error("Erro de conexão com o backend Go")
        return Response(
            {'error': 'Não foi possível conectar ao backend. Verifique se o servidor Go está rodando na porta 8080.'}, 
            status=status.HTTP_503_SERVICE_UNAVAILABLE
        )
    except Exception as e:
        logger.error(f"Erro inesperado: {str(e)}")
        return Response(
            {'error': 'Erro interno do servidor'}, 
            status=status.HTTP_500_INTERNAL_SERVER_ERROR
        )

@api_view(['GET'])
def health_check(request):
    """
    Endpoint para verificar a saúde do gateway e conectividade com o backend
    """
    try:
        # Testar conexão com o backend Go
        response = requests.get(f'{GO_BACKEND_URL}/livros', timeout=5)
        backend_status = response.status_code == 200
        
        return Response({
            'gateway': 'OK',
            'backend_go': 'OK' if backend_status else 'ERROR',
            'backend_url': GO_BACKEND_URL
        }, status=status.HTTP_200_OK)
        
    except requests.exceptions.ConnectionError:
        return Response({
            'gateway': 'OK',
            'backend_go': 'CONNECTION_ERROR',
            'backend_url': GO_BACKEND_URL,
            'message': 'Backend Go não está acessível'
        }, status=status.HTTP_200_OK)
    except Exception as e:
        return Response({
            'gateway': 'OK',
            'backend_go': 'ERROR',
            'backend_url': GO_BACKEND_URL,
            'error': str(e)
        }, status=status.HTTP_200_OK)
