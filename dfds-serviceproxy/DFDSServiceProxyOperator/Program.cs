using System;
using System.Threading.Tasks;
using k8s;

namespace DFDSServiceProxyOperator
{
    class Program
    {
        public static async Task Main(string[] args)
        {
            var clientConfig = KubernetesClientConfiguration.BuildDefaultConfig();
            var generic = new GenericClient(clientConfig, "stable.dfds.cloud", "v1alpha1", "servicesproxies");

            var crs = await generic.ListNamespacedAsync<CustomResourceList<ServiceProxy>>("developerautomation-xavgy");
            foreach (var cr in crs.Items)
            {
                Console.WriteLine("- CR Item {0} = {1}", crs.Items.IndexOf(cr), cr.Metadata.Name);
            }
        }
    }
}
