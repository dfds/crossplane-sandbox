XR:
    Composite resource

XRD:
    Definition of XR
    Cluster-wide
    Can optionally provide a XRC for namespaced XR

XRC:
    Claim of XR


articles, docs:
    - https://dev.to/sahillakhwani/building-custom-control-planes-using-crossplane-404f
    - https://crossplane.io/docs/v1.0/getting-started/compose-infrastructure.html
    - https://crossplane.io/docs/v0.14/getting-started/package-infrastructure.html


Thoughts:

- How to make services accessible:
  - Automatically create services on other clusters as soon as they're added to CENTRAL STORE
  - OR
  - Use the CRD(ServiceProxy) and let the user specify the services they need as the time
