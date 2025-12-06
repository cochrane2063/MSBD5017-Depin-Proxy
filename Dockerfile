FROM ghcr.io/wg-easy/wg-easy:14
WORKDIR /app
COPY start_wireguard.sh /
COPY clearnetNode /
RUN chmod +x /start_wireguard.sh /clearnetNode
CMD ["/start_wireguard.sh"]