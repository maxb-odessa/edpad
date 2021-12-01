
PREFIX		=	${HOME}/.local
SHAREDIR	=	${PREFIX}/share/edpad

GO111MODULE	=	auto
GOBIN		=	${PREFIX}/bin

all:
	go build ./cmd/edpad

install:
	go env -w GOBIN=${GOBIN}
	go install ./cmd/edpad
	mkdir -p ${SHAREDIR}
	cp -f resources/edpad.css.in ${SHAREDIR}/edpad.css
	sed -i "s=@SHAREDIR@=${SHAREDIR}=g" ${SHAREDIR}/edpad.css
	cp -f resources/edpad.glade ${SHAREDIR}
	cp -f resources/edpad-1280-800.png ${SHAREDIR}/edpad.png

