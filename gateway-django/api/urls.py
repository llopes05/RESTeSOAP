from django.urls import path
from . import views

urlpatterns = [
    path('livros/', views.livros_view, name='livros'),
    path('usuarios/', views.usuarios_view, name='usuarios'),
    path('emprestimos/', views.emprestimos_view, name='emprestimos'),
    path('health/', views.health_check, name='health_check'),
]
