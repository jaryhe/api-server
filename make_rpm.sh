#!/bin/bash

set -e
VERSION=`date +%-y`
while getopts "v:" opt; do
	case $opt in
	v)
        VERSION="$OPTARG"
	;;
    :)
         echo "Option -$OPTARG requires an argument." >&2
         exit 1
          ;;
    *)
          echo "Invalid option: -$OPTARG" >&2
          exit 1
          ;;
   esac
done

dir=`dirname $0`
SPEC_FILE=storm.spec

TOOLS_ROOT=`pwd`

BUILD="${TOOLS_ROOT}/build"

mkdir -p $BUILD/buildroot/{BUILD,RPMS,S{RPMS,PECS,OURCES}}
rm -f $BUILD/buildroot/SOURCES/*.tgz


tar czf $BUILD/buildroot/SOURCES/storm-$VERSION.tgz *

sed -e "s/%VERSION%/$VERSION/g" ${SPEC_FILE}.in > ${SPEC_FILE}
rpmbuild  -bb ${SPEC_FILE} --define "_topdir $BUILD/buildroot/" 

rm -rf ./dist
mkdir ./dist
cp -rf $BUILD/buildroot/RPMS/* ./dist

rm -rf $BUILD

