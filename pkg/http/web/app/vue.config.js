// vue.config.js
const path = require("path");

module.exports = {
  assetsDir: "assets",
  configureWebpack: {
    resolve: {
      alias: {
        "@src": path.resolve(__dirname, "src/"),
        "@components": path.resolve(__dirname, "src/components"),
        "@views": path.resolve(__dirname, "src/views"),
        "@layouts": path.resolve(__dirname, "src/layouts"),
        "@utils": path.resolve(__dirname, "src/utils"),
        "@assets": path.resolve(__dirname, "src/assets")
      }
    }
  },
  devServer: {
    proxy: "http://localhost:3000"
  }
};
