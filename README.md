# TPEGO

Para resolver el trabajo se creó una funcion llamada "Chain" donde se le pasa un string y su retorno será tendrá el tipo de la estructura creada llamada "Result".
Dentro de dicha función se detectan los siguientes cuatro campos del string: 
            tipo (porque type figuraba como palabra reservada), value y length; y tambien el largo del value. 
Los primeros tres valores se detectan con funciones creadas por el usuario, donde se les pasa el string de la funcion Chain y segun el valor a detectar se realizan
diferentes procedimientos. Mientras que para sacar el largo del value se utilizó la funcion len() de un paquete ya existente.

Una vez obtenidos los cuatro campos nombrados anteriormente se procede a hacer las comprobaciones para que el valor de retorno cumpla con la condicion de la consigna.
Primero que nada se comprueba que length sea mayor que cero, de ser asi, se comprueba que: 
    si el tipo es "NN" (number), que el value sea solo de numeros y que éste coincida con el length.
    si el tipo es "TX" (text), que el value sea solo de letras y que esté coincida con el lenth.
    
    

Para hacer el testing tuve que realizarlo sobre el archivo main.go, porque al querer realizarlo sobre el archivo chain.go que habia creado para separar las funciones del main 
me tiraba diferentes errores que no pude solucionarlos. Igualmente, cuando utilizaba las funciones en el archivo chain.go y en el main solo creaba y le pasaba la cadena me
funcionaba correctamente.
