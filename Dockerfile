FROM alpine:3.19
WORKDIR /app

# Copia o binário já compilado localmente
COPY main .
RUN chmod +x ./main

EXPOSE 8080
CMD ["./main"]