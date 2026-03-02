plugins {
    kotlin("jvm") version "1.9.0"
    id("maven-publish")
}

group = "com.sdkwork"
version = "1.0.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation("com.squareup.okhttp3:okhttp:4.12.0")
    implementation("com.fasterxml.jackson.core:jackson-databind:2.16.0")
    
    testImplementation("org.jetbrains.kotlin:kotlin-test")
}

publishing {
    publications {
        create<MavenPublication>("maven") {
            groupId = "com.sdkwork"
            artifactId = "sdk-common"
            version = "1.0.0"
            
            from(components["java"])
        }
    }
}
