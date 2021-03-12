module.exports = {
  devServer: {
    proxy: {
      "/api-capsvc": {
        target: 'https://api.hellman.oxygen.dfds.cloud',
        changeOrigin: true,
        pathRewrite: {
          '^/api-capsvc': '/capability/api'
        }
      },
      "/api": {
        target: 'http://localhost:8090',
        changeOrigin: true
      }
    }
  }
}
  