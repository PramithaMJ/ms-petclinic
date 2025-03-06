# Dockerfile
FROM node:14

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json /app
RUN npm install

# Copy the rest of the application code
COPY . /app

# Expose the port
EXPOSE 3000

# Start the application
CMD ["node", "index.js"]
