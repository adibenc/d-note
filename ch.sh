for i in $(seq 1 2); do echo $i;
echo "202010"$i" 09:17:00"
sudo date -s "2020-10-$i 09:17:00";
date;
echo $i" " | tee test;
git add .;
git commit -m "test "$i;
done;
