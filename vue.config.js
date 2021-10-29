module.exports = {
  configureWebpack: {
    devServer: {
      watchOptions: {
        poll: true
      },
      disableHostCheck: true
    }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
