FROM arm32v7/busybox:latest
EXPOSE 8000

COPY ./alertrack_rpi .
COPY ./alertrack.toml .

ENTRYPOINT [ "./alertrack_rpi" ]

