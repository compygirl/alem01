INTERVIEW=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $INTERVIEW
cat interviews/interview-"$INTERVIEW"
echo $MAIN_SUSPECT
