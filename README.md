Gopress will enable you to build exciting and modern presentations using impress.js by doing the thing that feels the most natural to us developers: coding.
Creating the content is as natural as writing markdown files, building a presentable version of your presentation is done by using a CLI.
Since everything is stored in code, you can even use version control to manage your presentations!

An example workflow looks like this:
1. Start by creating a directory for your presentations: mkdir -p my-awesome-slides && cd my-awsome-slides.
2. Create your first presentation: gopress create gopress-101.
3. Start writing your content using GitHub-Flavored markdown: echo '# GoPress\n## ...simply\n## ...works' > 00_intro.md.
4. Build your presentation and open it in your default browser: gopress present.