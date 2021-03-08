using Microsoft.Extensions.Options;
using Service.Classes;
using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Identity.Client;

namespace Service
{
    public class ServiceProxyService : IServiceProxyService
    {
        private readonly IOptions<ServiceProxySettings> _options;
        public ServiceProxyService(IOptions<ServiceProxySettings> options)
        {
            _options = options;
        }

        public async Task<List<ServiceProxyResult>> GetResults()
        {
            List<ServiceProxyResult> results = new List<ServiceProxyResult>();

            foreach (var p in _options.Value.proxyUrl)
            {
                var sp = new ServiceProxy(p, ConfidentialClientApplicationBuilder
                    .Create(_options.Value.clientId)
                    .WithClientSecret(_options.Value.clientSecret)
                    .WithTenantId("73a99466-ad05-4221-9f90-e7142aa2f6c1")
                    .Build());
                results.Add(await sp.GetResults());
            }

            return results;
        }

    }
}
