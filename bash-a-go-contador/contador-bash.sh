#!/bin/bash


# Comparaciones con numeros
# WORD=8
# if [ $WORD -eq 10 ]; then #si son igual WORD=8 y el valor 10
# if [ $WORD -ne 10 ]; then #si no son igual
# if [ $WORD -lt 10 ]; then #si es menor
# if [ $WORD -gt 10 ]; then #si es mayor
# if [ $WORD -le 10 ]; then #si es menor o igual
# if [ $WORD -ge 5 ]; then #si es mayor o igual
#     echo "Run is fun!!"
# else
#     echo "--------"

# fi

# Comparaciones strings
# WORD="run"
# if [ $WORD == "run" ]; then
#     echo "Run is fun!!"
# else
#     echo "--------"

# fi
# Comparaciones strings2

# WORD_TO_TEST="runtest"
# if [ $WORD_TO_TEST == "runtest" ] -o [ $WORD == "run" ] ; then
# echo "Running is fun!!"
# else
# echo "--------"
# fi
# Comparaciones strings3

WORD3="run33"
WORD_TO_TEST3="runtest33"
WORD4="run43"
WORD_TO_TEST4="runtest4"



# <<<<<<< HEAD
# =======
if [ $WORD_TO_TEST3 == "runtest3" ] || [ $WORD3 == "run3" ] || [ $WORD4 == "run4" ]|| [ $WORD_TO_TEST4 == "runtest4" ]; then
# >>>>>>> 252a6804bf110972304341cf777ccc7d03c173fc
echo "Running is fun!!"
else
echo "--------"
fi


# Ejemplo1
# while true; do
#     echo "Presiona Cotrl + C para salir"
#     sleep 3s
# done


# WORD="run
# echo $WORD
# echo "${WORD}ning is fun!"
# echo '${WORD}ning is fun!'
# echo "===================="

# Ejemplo2
CONTADOR=0

#Cuando sea menor o igual
while [ $CONTADOR -lt 6 ]; do
    if [ $CONTADOR -eq 0 ]; then
        echo "El contador es cero"
    else
        echo "El contador es diferente a cero"
    fi
    CONTADOR=$(($CONTADOR + 1))
    sleep 3s
done