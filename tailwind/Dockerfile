FROM node:latest
WORKDIR /tailwind
RUN npm init -y && \
    npm install tailwindcss@^3.3.3

CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll