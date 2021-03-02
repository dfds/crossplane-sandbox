using Microsoft.Extensions.DependencyInjection;
using System;
using System.Collections.Generic;
using System.Text;

namespace Service
{
    public static class IServiceCollectionExtension
    {
        public static IServiceCollection AddServiceProxyServiceCollection(this IServiceCollection services, string proxyUrl)
        {
            services.Configure<ServiceProxySettings>(
                options =>
                {
                    options.proxyUrl = proxyUrl;
                });

            services.AddTransient<IServiceProxy, ServiceProxy>();
            return services;
        }
    }
}
