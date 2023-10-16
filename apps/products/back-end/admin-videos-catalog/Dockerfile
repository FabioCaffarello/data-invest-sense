FROM docker.io/node:lts-alpine

ENV HOST=0.0.0.0
ENV PORT=3000

WORKDIR /app

RUN addgroup --system products-back-end-admin-videos-catalog && \
          adduser --system -G products-back-end-admin-videos-catalog products-back-end-admin-videos-catalog

COPY dist/apps/products/back-end/admin-videos-catalog products-back-end-admin-videos-catalog
RUN chown -R products-back-end-admin-videos-catalog:products-back-end-admin-videos-catalog .

COPY package.json .

# You can remove this install step if you build with `--bundle` option.
# The bundled output will include external dependencies.
RUN npm install --force

CMD [ "node", "products-back-end-admin-videos-catalog" ]
