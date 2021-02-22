using k8s.Models;
using Newtonsoft.Json;

namespace DFDSServiceProxyOperator
{
    public class ServiceProxy : CustomResource<ServiceProxySpec, ServiceProxyStatus>
    {
        public override string ToString()
        {
            var labels = "{";
            foreach (var kvp in Metadata.Labels)
            {
                labels += kvp.Key + " : " + kvp.Value + ", ";
            }

            labels = labels.TrimEnd(',', ' ') + "}";

            return $"{Metadata.Name} (Labels: {labels}), Spec.Services: {Spec.Properties.Services}";
        }
    }

    public class ServiceProxySpec
    {
        [JsonProperty(PropertyName = "properties")]
        public ServiceProxySpecProperties Properties { get; set; }
    }

    public class ServiceProxySpecProperties
    {
        [JsonProperty(PropertyName = "services", NullValueHandling = NullValueHandling.Ignore)]
        public string Services { get; set; }
    }

    public class ServiceProxyStatus : V1Status
    {
        
    }
}