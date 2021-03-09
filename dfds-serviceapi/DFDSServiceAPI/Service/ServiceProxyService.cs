﻿using Microsoft.Extensions.Options;
using Service.Classes;
using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

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
                var sp = new ServiceProxy(p.ToString());
                results.Add(await sp.GetResults());
            }

            return results;
        }

    }
}