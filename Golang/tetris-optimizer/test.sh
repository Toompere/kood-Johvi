echo "bad example 00:"
go run . test/badexample00.txt
sleep 2
echo
echo "bad example 01:"
go run . test/badexample01.txt
sleep 2
echo
echo "bad example 02:"
go run . test/badexample02.txt
sleep 2
echo
echo "bad example 03:"
go run . test/badexample03.txt
sleep 2
echo
echo "bad example 04:"
go run . test/badexample04.txt
sleep 2
echo
echo "bad format:"
go run . test/badformat.txt
sleep 2
echo
echo "good example 00:"
go run . test/goodexample00.txt
sleep 2
echo
echo "good example 01:"
go run . test/goodexample01.txt
sleep 2
echo
echo "good example 02:"
go run . test/goodexample02.txt
sleep 2
echo
echo "good example 03:"
go run . test/goodexample03.txt
sleep 2
echo
echo "hard example:"
echo "This is going to take some time!"
go run . test/hardexample.txt