# Readme

Irregularly updated blog about random things. See at https://ibnishak.github.io/blog/

Feel free to start issues if you have any comments.

## Setup and installation of dependecies
- [Connect github via ssh](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)
- Switch remote url to ssh if you have not already
```zsh
git remote set-url origin git@github.com:ibnishak/blog.git
```
- Run the `./install.zsh`




## Digital garden

- See the pages here: https://ibnishak.github.io/digital-garden/
- Made with [gatsby-digital-garden](https://github.com/mathieudutour/gatsby-digital-garden)

# Staring a new blog

1. Shallow clone `main` branch of this repo.

```
git clone -b main --single-branch --depth=1 git@github.com:ibnishak/digital-garden.git
```

2. Change address of origin remote.

```
git remote set-url "my-new-url"
```

3. Open `gatsby-config.js` and update the options like path-prefix, site title, description.
4. Run `npm install`
5. Add markdown/mdx files to `content` folder
6. Run `npm run deploy`

# Updating existing blog

1. Clone the repo.
2. Run `npm install`
3. Add markdown/mdx files to `content` folder
4. Run `npm run deploy`