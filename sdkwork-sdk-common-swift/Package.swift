// swift-tools-version:5.7
import PackageDescription

let package = Package(
    name: "SDKworkCommon",
    platforms: [
        .iOS(.v13),
        .macOS(.v10_15),
    ],
    products: [
        .library(
            name: "SDKworkCommon",
            targets: ["SDKworkCommon"]
        ),
    ],
    dependencies: [],
    targets: [
        .target(
            name: "SDKworkCommon",
            dependencies: [],
            path: "Sources"
        ),
        .testTarget(
            name: "SDKworkCommonTests",
            dependencies: ["SDKworkCommon"],
            path: "Tests"
        ),
    ]
)
