UPDATE users SET balance=balance-$1 WHERE email=$2 and balance>=$1;
