#!/usr/bin/env bash
# exit when any command fails
set -e


DOMAINTYPESPATH=github.com/theNorstroesm/todo-specs/dist/pb/
TARGETDIR="../pb"
# enable recursion for /**/*.xxx
shopt -s globstar dotglob
WD=`pwd`

cd dist/protos/

TMPDIR=$TARGETDIR"_tmp"
rm -rf $TMPDIR
mkdir $TMPDIR

FILES=./**/*.proto


FILES=./**/*.proto
for f in $FILES
do
  #echo "Processing $f file..."

# the following lines are included, because the furo types are not complete :-(
# Mgoogle/type/date.proto=google.golang.org/genproto/genproto/googleapis/type/date,\
# Mgoogle/type/money.proto=google.golang.org/genproto/genproto/googleapis/type/money,\


protoc --proto_path=./ \
-I. \
-I$WD/dependencies/github.com/theNorstroem/furoBaseSpecs/dist/proto \
-I/usr/local/include \
--go-grpc_out=\
:$TMPDIR \
--go_out=\
:$TMPDIR \
--grpc-gateway_out=logtostderr=true:$TMPDIR  $f


done


cd $TMPDIR/github.com/theNorstroesm/todo-specs

FILES=**/*.go


for f in $FILES
do
dir=$(dirname -- "$WD/$f")
mkdir -p $dir

file=$WD/$f
[ -f $file ] && rm $file
mv $f $WD/$f
done

cd $WD/dist/protos/
rm -rf $TMPDIR
