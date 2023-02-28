using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace skillathon_api.Controllers
{
    [ApiController]
    [Route("/")]
    public class skillathonController : ControllerBase
    {  [HttpGet]
        public IActionResult Get()
        {
           return Ok("Deployment Done !!");
        }
    }
}
