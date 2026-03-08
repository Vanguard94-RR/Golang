# Ejercicios - Sección 3: Values, Variables y Constantes

## Ejercicio 1: Conversión de Temperaturas

**Descripción:**
Crea un programa que convierta una temperatura de Celsius a Fahrenheit. Usa constantes para los valores de conversión (C a F: multiplica por 9/5 y suma 32).

**Entrada y Salida Esperada:**
```
Celsius: 0°C → Fahrenheit: 32°F
Celsius: 25°C → Fahrenheit: 77°F
Celsius: 100°C → Fahrenheit: 212°F
```

**Pista:**
Define la fórmula como constantes. Declara las temperaturas como variables e imprime los resultados.

---

## Ejercicio 2: Cálculo de IMC (Índice de Masa Corporal)

**Descripción:**
Programa que calcule el IMC dado peso (kg) y altura (m). La fórmula es: IMC = peso / (altura²).

**Entrada y Salida Esperada:**
```
Peso: 70 kg, Altura: 1.75 m → IMC: 22.86
Peso: 80 kg, Altura: 1.80 m → IMC: 24.69
```

**Pista:**
Declara variables para peso y altura. Usa la función `math.Pow()` para calcular el cuadrado, o simplemente multiplica altura por sí misma.

---

## Ejercicio 3: Contador de Segundos en Período de Tiempo

**Descripción:**
Crea un programa que calcule cuántos segundos hay en un período de tiempo dado (horas, minutos, segundos). Usa constantes para las conversiones.

**Entrada y Salida Esperada:**
```
2 horas, 30 minutos, 45 segundos → 9045 segundos totales
1 hora, 0 minutos, 0 segundos → 3600 segundos totales
0 horas, 5 minutos, 30 segundos → 330 segundos totales
```

**Pista:**
Define constantes: SecondsPerMinute = 60, SecondsPerHour = 3600. Declara variables para cada unidad y suma los totales.
