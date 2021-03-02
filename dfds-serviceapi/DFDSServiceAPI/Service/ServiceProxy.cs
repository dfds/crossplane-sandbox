using Microsoft.Extensions.Options;
using Service.Classes;
using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace Service
{
    public class ServiceProxy : IServiceProxy
    {
        private readonly HttpClient _client = new HttpClient();

        public ServiceProxy(IOptions<ServiceProxySettings> options)
        {
            _client.BaseAddress = new Uri(options.Value.proxyUrl);
        }

        public async Task<ServiceProxyResult> GetResults()
        {
            ServiceProxyResult result = new ServiceProxyResult();
            using (_client)
            {
                result.result = await _client.GetAsync("/api/get-all").Result.Content.ReadAsStringAsync();
            }
            return result;
        }
    }
}
