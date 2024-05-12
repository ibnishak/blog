module.exports = {
 pathPrefix: "/blog",
  siteMetadata: {
    title: `DryMind`,
    description: `Thoughts, notes and views`,
  },
  plugins: [
    {
      resolve: `gatsby-theme-garden`,
      options: {
	contentPath: `${__dirname}/gatsbycontent`,
        basePath: `/`,
        rootNote: `/About`,
        parseWikiLinks: true,
      },
    },
  ],
};
