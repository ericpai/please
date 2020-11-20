net use \\%1 "%2" /user:"%3"
echo "F" | xcopy /y \\%1\%4 %5
@REM net use \\10.0.0.13 "developer" /user:"developer"
@REM echo "F" | xcopy /y \\10.0.0.13\c$\test.txt D:\test.txt: