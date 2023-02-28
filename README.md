# cloud-devops-skillathon 2023

To run projects. Go into directory and type:

skillathon-api - `dotnet run`

skillathon-ui - `npm run start`

Azure Subscription and Resource Group - `https://portal.azure.com/#@ShellCorp.onmicrosoft.com/resource/subscriptions/94d67791-6e06-4f77-bc46-0b862e8aeb03/resourceGroups/sbox-skillathonq1-dev-rg/overview`

Azure DevOps - `https://sede-sandbox.visualstudio.com/IDESandbox`

Service Connection - `IDESandbox (GitHub) and IDESandBoxAzure (Azure)`

<hr/>
<h1><center>Challenge</center></h1> 

<li>Create separate resources for API and UI in Azure.</li> 
<li>Build and Release pipelines creation for resources.</li>
<li>Deployment should be done from master(default) branch only.</li>
<li>No direct check in code to master(default) branch.</li>
<li>Automatic build and deployment after code pushed to master (default) branch.</li>
<li>Use Classic Pipelines for ADO.</li> 

<br/>
<h2>Bonus !</h2>
<li>How many environments and resources?</li>
<li>Can you build and release from single pipeline for multiple applications in this case UI and API?</li>
<li>How do you get approval of the stakeholders before deployment to production?</li>
<li>Should Dependent API and UI be dependant deployments? what should be the execution flow?</li>
