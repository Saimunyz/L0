# L0
Тестовое задание - L0  

## Для запуска  
1) Нужно запустить docker-compose up -d для запуска postgres и nats-streaming
2) Запустить скрипт для отправки данных через nats-streaming: make data
3) Запустить проект make run

После запуска make run нужно открыть в браузере http://127.0.0.1:8080/  
Вся основная кофигурация проекта лежит в config.yaml. Данные находятся в data.json

## WRK тесты
    Running 10s test @ http://127.0.0.1:8080/orders/4e85a1fb33a2233dragon  
    12 threads and 400 connections
|Thread Stats | Avg   | Stdev     | Max      | +/- Stdev|  
|-------------|-------|-----------|----------|----------|   
|Latency      | 283.20ms| 284.96ms   | 1.98s |  86.47% |
|Req/Sec     |137.00    |48.71   | 383.00   |70.08%   |
    16368 requests in 10.09s, 75.74MB read 
    Socket errors: connect 0, read 0, write 0, timeout 1
    Requests/sec:   1621.80  
    Transfer/sec:      7.50MB  
--------------------------------------------------------
    Running 10s test @ http://127.0.0.1:8080/
    12 threads and 400 connections
|Thread Stats   |Avg      |Stdev    | Max  | +/- Stdev|  
| ---------- | ---------| --------| -------| ---------|
|Latency   |315.47ms | 321.46ms |  1.99s  |  81.52% |
|Req/Sec  | 126.31   |  44.49  | 435.00   |  69.62% |
    15119 requests in 10.09s, 19.84MB
    Socket errors: connect 0, read 0, write 0, timeout 5  
    Requests/sec:   1497.94  
    Transfer/sec:      1.97MB  
