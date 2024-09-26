FROM node:20-alpine as builder
WORKDIR /usr/src/app
COPY ./package.json ./
RUN yarn install --frozen-lockfile
# COPY ./.env.example ./.env.local
COPY . .
RUN yarn build

FROM node:20-alpine as runner
WORKDIR /usr/src/app
ENV NODE_ENV production
ENV PORT 3002
COPY --from=builder /usr/src/app/next.config.mjs ./
COPY --from=builder /usr/src/app/.next/standalone ./
COPY --from=builder /usr/src/app/.next/static ./.next/static
# COPY --from=builder /usr/src/app/.env.local ./.env.local

CMD ["node", "server.js"]