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

protoc --proto_path=./ \
-I./ \
-I$WD/dependencies/github.com/theNorstroem/furoBaseSpecs/dist/proto \
-I/usr/local/include \
--go_out=\
:$TMPDIR \
--go-grpc_out=\
:$TMPDIR $FILES

cd $TMPDIR/github.com/theNorstroesm/todo-specs

FILES=**/*.go


for f in $FILES
do
dir=$(dirname -- "$WD/$f")
# echo $f
mkdir -p $dir
file=$WD/$f
[ -f $file ] && rm $file
mv $f $WD/$f
done

cd $WD/dist/protos/
pwd
rm -rf $TMPDIR

