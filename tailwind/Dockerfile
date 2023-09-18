FROM node:latest
WORKDIR /tailwind
RUN npm init -y && \
    npm install tailwindcss && \
    npx tailwindcss init
CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /src/styles.css --watch --poll


