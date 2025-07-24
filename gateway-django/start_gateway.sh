#!/bin/bash

# Script para executar o gateway Django

echo "🚀 Iniciando Gateway Django para Sistema de Biblioteca"
echo "==============================================="

# Verificar se o ambiente virtual existe
if [ ! -d "venv" ]; then
    echo "❌ Ambiente virtual não encontrado!"
    echo "Execute: python -m venv venv"
    exit 1
fi

# Ativar ambiente virtual
source venv/Scripts/activate

# Verificar se as dependências estão instaladas
if ! python -c "import django" 2>/dev/null; then
    echo "📦 Instalando dependências..."
    pip install -r requirements.txt
fi

# Fazer migrações (caso necessário)
echo "🔄 Verificando migrações..."
python manage.py makemigrations
python manage.py migrate

# Iniciar servidor
echo "🌐 Iniciando servidor na porta 8000..."
echo "📚 Gateway URL: http://localhost:8000/"
echo "🔍 Health Check: http://localhost:8000/api/health/"
echo "📋 API Endpoints:"
echo "   - GET/POST http://localhost:8000/api/livros/"
echo "   - GET/POST http://localhost:8000/api/usuarios/"
echo "   - GET/POST http://localhost:8000/api/emprestimos/"
echo ""
echo "⚡ IMPORTANTE: Certifique-se de que o backend Go está rodando na porta 8080!"
echo ""

python manage.py runserver 0.0.0.0:8000
