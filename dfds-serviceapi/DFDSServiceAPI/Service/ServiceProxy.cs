using k8s;
using k8s.Models;
using Microsoft.Extensions.Options;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using Service.Classes;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace Service
{
    public class ServiceProxy : IServiceProxy
    {
        private readonly HttpClient _client = new HttpClient();

        public ServiceProxy(string proxyUrl)
        {
            _client.BaseAddress = new Uri(proxyUrl);
        }

        public async Task<ServiceProxyResult> GetResults()
        {
            ServiceProxyResult result = new ServiceProxyResult(_client.BaseAddress.ToString());

            using (_client)
            {     
                _client.BaseAddress = new Uri(_client.BaseAddress.ToString());

                try
                {
                    var temp = await _client.GetAsync("/api/get-all").Result.Content.ReadAsStringAsync();
                    var json = JObject.Parse(temp);
                    var ingress = json["Ingress"];
                    var service = json["Service"];

                    foreach (var x in ingress)
                    {
                        var s = JsonConvert.DeserializeObject<Extensionsv1beta1Ingress>(x.ToString());
                        result.ingresses.Add(s);
                    }

                    foreach (var y in service)
                    {
                        var s = JsonConvert.DeserializeObject<V1APIService>(y.ToString());
                        result.services.Add(s);
                    }
                }
                catch (Exception e)
                {
                    Console.WriteLine(e);
                }
            }

            return result;
        }
    }
}
