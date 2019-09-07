FROM scratch
ADD dist/entry-point /entry-point
ENTRYPOINT [ "/entry-point" ]