# --- Stage 1: Install Dependencies ---
FROM node:20-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /web

COPY package.json package-lock.json* ./
RUN npm ci

# --- Stage 2: Build the Application ---
FROM node:20-alpine AS builder
WORKDIR /web
COPY --from=deps /web/node_modules ./node_modules
COPY . .

# Environment variables needed at build time can be passed here
# ENV NEXT_PUBLIC_API_URL=http://localhost:8080

RUN npm run build

# --- Stage 3: Production Runner ---
FROM node:20-alpine AS runner
WORKDIR /web

ENV NODE_ENV=production
ENV PORT=3000
ENV HOSTNAME="0.0.0.0"

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# Automatically leverage output traces to reduce image size
COPY --from=builder /web/public ./public
COPY --from=builder --chown=nextjs:nodejs /web/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /web/.next/static ./.next/static

USER nextjs

EXPOSE 3000

CMD ["node", "server.js"]