# GoGoProjectTemplates

GoGoProjectTemplates is a collection of Go project templates designed to streamline the setup of new Go projects. Each
folder in the repository root represents an individual Go module that can be used as a standalone template with [
`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew).

## Important Note

The **base repository** itself is **not intended to be used with `gonew`**. Instead, only the individual folders within
the repository root (each containing its own Go module) are designed to be used as templates. If you want to contribute
or customize the templates, you can clone this repository. Otherwise, use `gonew` to initialize new projects based on
the specific template folders.

## Features

- **Template Structure:** Each folder in the repository is a self-contained Go module designed for a specific purpose (
  e.g., CLI application, microservice).
- **Optimized for `gonew`:** Quickly initialize new projects by selecting the desired template folder.
- **Customizable:** Clone the repository to modify or extend the templates for specific use cases.
- **Best Practices Included:** Each template follows Go development standards, including dependency management and
  project organization.

## Requirements

- Go (version 1.23 or later)
- [`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) installed

### Install `gonew`

To install `gonew`, run the following command:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

## Using the Templates

1. Choose the template folder that suits your project (e.g., `cli-template`, `microservice-template`).
2. Use `gonew` to initialize a new project based on that template:

   ```bash
   gonew init <module_name> GoGoProjectTemplates/<template_name> [options]
   ```

   Replace:
    - `<module_name>` with the desired name of your new Go module.
    - `<template_name>` with the name of the folder/template in the repository root (e.g., `cli-template`).

3. Navigate to your new project:

   ```bash
   cd <module_name>
   ```

4. Start building your project based on the selected template!

## Customization and Contribution

If you'd like to customize the templates or contribute new ones, you can clone this repository:

```bash
git clone <repository_url>
```

Inside the repository, each folder is its own Go module. You can:

- Update existing templates
- Add new templates as new folders
- Test your changes locally before sharing

Once you've made improvements, feel free to submit a pull request to contribute back to the project.

## Contributing

We appreciate contributions that improve or expand the available templates! To contribute:

1. Fork the repository.
2. Add or modify templates within the repository.
3. Open a pull request with clear changes and descriptions.

## License

GoGoProjectTemplates is licensed under the [GPL-3.0 License](LICENSE).
