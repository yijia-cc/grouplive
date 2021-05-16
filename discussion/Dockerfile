# Build
FROM openjdk:11.0.2-jdk-oracle AS builder

WORKDIR /app

COPY . .

RUN ./gradlew build

# Production
FROM openjdk:11.0.2-jre-slim AS production

WORKDIR /app

#RUN apk add --no-cache bash

COPY --from=builder /app/build/libs/*.jar app.jar

ENTRYPOINT ["java", "-jar", "app.jar"]
