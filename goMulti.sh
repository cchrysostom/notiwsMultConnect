#/bin/bash
LOOPS=5
if [ $1 ]
  then
    LOOPS=$1
fi

COUNT=0
while [ $COUNT -lt $LOOPS ]
  do
    echo Iteration, $COUNT.
    ./multi
    COUNT=$[$COUNT + 1]
done
