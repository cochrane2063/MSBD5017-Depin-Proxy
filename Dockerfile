FROM ghcr.io/wg-easy/wg-easy:14
# FROM alpine
WORKDIR /app
COPY start_wireguard.sh /start_wireguard.sh
COPY clearnetNodeJodAlpine /app/clearnetNodeJodAlpine
COPY .env /app/.env
RUN chmod +x /start_wireguard.sh /app/clearnetNodeJodAlpine
CMD ["/start_wireguard.sh"]
# CMD ["/app/clearnetNodeJodAlpine"]