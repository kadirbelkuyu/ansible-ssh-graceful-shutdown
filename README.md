# Ansible SSH Graceful Shutdown

This project is an example of using Ansible and SSH to control multiple target systems. The project can be used to perform a graceful shutdown of the target systems.

## Usage

1. Navigate to the project directory.
2. Open the `inventory` file and add the IP addresses or DNS names of the target systems.
3. Open the `main.go` file and make the necessary changes (username, password, etc.).
4. Run the command `go run main.go` on the command line.
5. The program will initiate graceful shutdown on the target systems.

## Requirements

- Go (https://golang.org/)
- Target systems with SSH access

## Important Note

- This example works with a user and password with SSH access to the systems. For security reasons, we recommend using SSH key-based authentication.
- This example can be used to perform graceful shutdown on the target systems. However, please ensure all important data is backed up before performing this operation.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
