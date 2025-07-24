@echo off
REM Script para executar o gateway Django no Windows

echo 🚀 Iniciando Gateway Django para Sistema de Biblioteca
echo ===============================================

REM Verificar se o ambiente virtual existe
if not exist "venv\" (
    echo ❌ Ambiente virtual não encontrado!
    echo Execute: python -m venv venv
    pause
    exit /b 1
)

REM Ativar ambiente virtual
call venv\Scripts\activate.bat

REM Verificar se as dependências estão instaladas
python -c "import django" 2>nul
if errorlevel 1 (
    echo 📦 Instalando dependências...
    pip install -r requirements.txt
)

REM Fazer migrações (caso necessário)
echo 🔄 Verificando migrações...
python manage.py makemigrations
python manage.py migrate

REM Iniciar servidor
echo 🌐 Iniciando servidor na porta 8000...
echo 📚 Gateway URL: http://localhost:8000/
echo 🔍 Health Check: http://localhost:8000/api/health/
echo 📋 API Endpoints:
echo    - GET/POST http://localhost:8000/api/livros/
echo    - GET/POST http://localhost:8000/api/usuarios/
echo    - GET/POST http://localhost:8000/api/emprestimos/
echo.
echo ⚡ IMPORTANTE: Certifique-se de que o backend Go está rodando na porta 8080!
echo.

python manage.py runserver 0.0.0.0:8000
