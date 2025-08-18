## 📖 What I studied today:
- Rust Topic: Introduction to Rust
- Installation, Cargo and creating the first Rust project

## 🧠 Key Learnings:
- Installed Rust using Curl
- created a Rust project using Cargo
- Cargo is a Rust package manager and build system( it helps you develop and manage your packages).   Packages are other programs that you add to your project.
- The `cargo new <project name>` helps you to create a new Rust project. 
- To run your project, you can either ``cargo build --> ./target/<program name>`` or run ``cargo run``, which effectively builds and runs your program. 
- Adding ``--release`` to build generates your optimised release program for distribution.
- You can use the ``cargo check`` command to check if your code can be compiled without running the program.
- You can also use the ``cargo init`` command to generate your project in an existing folder.
- When you use the `cargo new `command, it creates a folder structure in the following diagram:

 
 ```mermaid
flowchart TD
Root[Project name] --contains your project files --> src
src --a must-have entry file--> main.rs
Root[Project name] --a hidden folder--> target 
target --contains your debug build--> debug
target --contains your release build--> release
Root[Project name] --> Cargo.toml
Root[Project name] --> Cargo.lock
 ```

## 🧩 Problems I Faced:
- none

## 🚀 Breakthroughs and Solutions:
- 

## 📓 Devlog Ideas:
- none

## 🎯 Tomorrow's Focus:
- Programming the guess game in rust.