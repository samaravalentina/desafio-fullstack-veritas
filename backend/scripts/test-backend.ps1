# Script de Teste do Backend - Kanban API
# PowerShell Script para testar todos os endpoints

Write-Host "=== Teste do Backend - Kanban API ===" -ForegroundColor Green
Write-Host ""

$baseUrl = "http://localhost:8080"
$headers = @{"Content-Type" = "application/json"}

# Aguardar servidor iniciar
Write-Host "Aguardando servidor iniciar..." -ForegroundColor Yellow
Start-Sleep -Seconds 3

# Teste 1: GET /tasks - Listar tarefas (deve retornar array vazio)
Write-Host "`n1. Testando GET /tasks (Listar tarefas)" -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method GET
    Write-Host "   ✓ Sucesso! Retornou: $($response | ConvertTo-Json)" -ForegroundColor Green
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
    exit 1
}

# Teste 2: POST /tasks - Criar tarefa 1
Write-Host "`n2. Testando POST /tasks (Criar tarefa 1)" -ForegroundColor Cyan
$task1 = @{
    title = "Tarefa de Teste 1"
    description = "Esta é uma tarefa de teste"
    status = "todo"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method POST -Headers $headers -Body $task1
    $task1Id = $response.id
    Write-Host "   ✓ Tarefa criada com ID: $task1Id" -ForegroundColor Green
    Write-Host "   Título: $($response.title)" -ForegroundColor Gray
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
    exit 1
}

# Teste 3: POST /tasks - Criar tarefa 2
Write-Host "`n3. Testando POST /tasks (Criar tarefa 2)" -ForegroundColor Cyan
$task2 = @{
    title = "Tarefa de Teste 2"
    description = "Outra tarefa de teste"
    status = "in_progress"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method POST -Headers $headers -Body $task2
    $task2Id = $response.id
    Write-Host "   ✓ Tarefa criada com ID: $task2Id" -ForegroundColor Green
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
}

# Teste 4: GET /tasks - Listar todas as tarefas
Write-Host "`n4. Testando GET /tasks (Listar todas as tarefas)" -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method GET
    Write-Host "   ✓ Total de tarefas: $($response.Count)" -ForegroundColor Green
    foreach ($task in $response) {
        Write-Host "   - ID: $($task.id), Título: $($task.title), Status: $($task.status)" -ForegroundColor Gray
    }
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
}

# Teste 5: PUT /tasks/:id - Atualizar tarefa
Write-Host "`n5. Testando PUT /tasks/$task1Id (Atualizar tarefa)" -ForegroundColor Cyan
$updatedTask = @{
    title = "Tarefa Atualizada"
    description = "Descrição atualizada"
    status = "done"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks/$task1Id" -Method PUT -Headers $headers -Body $updatedTask
    Write-Host "   ✓ Tarefa atualizada!" -ForegroundColor Green
    Write-Host "   Novo título: $($response.title)" -ForegroundColor Gray
    Write-Host "   Novo status: $($response.status)" -ForegroundColor Gray
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
}

# Teste 6: GET /tasks - Verificar atualização
Write-Host "`n6. Testando GET /tasks (Verificar atualização)" -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method GET
    $updatedTaskObj = $response | Where-Object { $_.id -eq $task1Id }
    if ($updatedTaskObj.status -eq "done") {
        Write-Host "   ✓ Tarefa atualizada corretamente!" -ForegroundColor Green
    } else {
        Write-Host "   ✗ Tarefa não foi atualizada corretamente" -ForegroundColor Red
    }
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
}

# Teste 7: DELETE /tasks/:id - Excluir tarefa
Write-Host "`n7. Testando DELETE /tasks/$task2Id (Excluir tarefa)" -ForegroundColor Cyan
try {
    $response = Invoke-WebRequest -Uri "$baseUrl/tasks/$task2Id" -Method DELETE
    if ($response.StatusCode -eq 204) {
        Write-Host "   ✓ Tarefa excluída com sucesso!" -ForegroundColor Green
    } else {
        Write-Host "   ✗ Status code inesperado: $($response.StatusCode)" -ForegroundColor Red
    }
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
}

# Teste 8: GET /tasks - Verificar exclusão
Write-Host "`n8. Testando GET /tasks (Verificar exclusão)" -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/tasks" -Method GET
    $remainingTasks = $response | Where-Object { $_.id -eq $task2Id }
    if ($remainingTasks.Count -eq 0) {
        Write-Host "   ✓ Tarefa excluída corretamente!" -ForegroundColor Green
        Write-Host "   Tarefas restantes: $($response.Count)" -ForegroundColor Gray
    } else {
        Write-Host "   ✗ Tarefa ainda existe" -ForegroundColor Red
    }
} catch {
    Write-Host "   ✗ Erro: $_" -ForegroundColor Red
}

# Teste 9: Validação - Título vazio
Write-Host "`n9. Testando validação (Título vazio)" -ForegroundColor Cyan
$invalidTask = @{
    title = ""
    description = "Tarefa sem título"
    status = "todo"
} | ConvertTo-Json

try {
    $response = Invoke-WebRequest -Uri "$baseUrl/tasks" -Method POST -Headers $headers -Body $invalidTask -ErrorAction Stop
    Write-Host "   ✗ Deveria ter retornado erro!" -ForegroundColor Red
} catch {
    if ($_.Exception.Response.StatusCode -eq 400) {
        Write-Host "   ✓ Validação funcionando corretamente!" -ForegroundColor Green
    } else {
        Write-Host "   ✗ Erro inesperado: $_" -ForegroundColor Red
    }
}

# Teste 10: Validação - Status inválido
Write-Host "`n10. Testando validação (Status inválido)" -ForegroundColor Cyan
$invalidStatusTask = @{
    title = "Tarefa com status inválido"
    description = "Teste"
    status = "invalid_status"
} | ConvertTo-Json

try {
    $response = Invoke-WebRequest -Uri "$baseUrl/tasks" -Method POST -Headers $headers -Body $invalidStatusTask -ErrorAction Stop
    Write-Host "   ✗ Deveria ter retornado erro!" -ForegroundColor Red
} catch {
    if ($_.Exception.Response.StatusCode -eq 400) {
        Write-Host "   ✓ Validação de status funcionando!" -ForegroundColor Green
    } else {
        Write-Host "   ✗ Erro inesperado: $_" -ForegroundColor Red
    }
}

Write-Host "`n=== Testes Concluídos ===" -ForegroundColor Green
Write-Host "`nVerifique o arquivo backend/tasks.json para confirmar persistência" -ForegroundColor Yellow



