using AutoMapper;
using DFDSServiceAPI.Dtos;
using Microsoft.AspNetCore.Mvc;
using Service;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace DFDSServiceAPI.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ServiceProxyController : Controller
    {
        private readonly IServiceProxy _proxy;
        private readonly IMapper _mapper;

        public ServiceProxyController(IServiceProxy proxy, IMapper mapper)
        {
            _proxy = proxy;
            _mapper = mapper;
        }

        [HttpGet]
        public async Task<IActionResult> GetAll()
        {
            var results = _mapper.Map<ServiceProxyResultDto>(await _proxy.GetResults());
            return Ok(results);
        }
    }
}
