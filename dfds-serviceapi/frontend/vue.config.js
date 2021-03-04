module.exports = {
    pluginOptions: {
      proxy: {
        enabled: true,
        context: '/api',
        options: {
          target: 'http://localhost:8090',
          changeOrigin: true
        }
      }
    }
  }
  