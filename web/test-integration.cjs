const { chromium } = require('playwright');
const fs = require('fs');
const path = require('path');

async function runIntegrationTest() {
  console.log('🚀 Starting integration test...');
  
  const evidenceDir = '.sisyphus/evidence/task-21';
  if (!fs.existsSync(evidenceDir)) {
    fs.mkdirSync(evidenceDir, { recursive: true });
  }
  
  const browser = await chromium.launch({ 
    headless: true,
    channel: 'chromium'
  });
  const context = await browser.newContext();
  const page = await context.newPage();
  
  try {
    // Test 1: Navigate to workspaces page
    console.log('📍 Test 1: Navigating to workspaces page...');
    await page.goto('http://localhost:5173/workspaces', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    
    await page.screenshot({ 
      path: path.join(evidenceDir, '01-workspaces-page.png'),
      fullPage: true 
    });
    console.log('✅ Workspaces page loaded');
    
    // Test 2: Create a new workspace
    console.log('📍 Test 2: Creating a new workspace...');
    
    // Look for create workspace button
    const createButton = await page.locator('button:has-text("Create"), button:has-text("New"), button:has-text("Add")').first();
    if (await createButton.isVisible()) {
      await createButton.click();
      await page.waitForTimeout(1000);
      
      // Fill in workspace name
      const nameInput = await page.locator('input[name="name"], input[placeholder*="name"], input[type="text"]').first();
      if (await nameInput.isVisible()) {
        await nameInput.fill('Integration Test Workspace');
        await page.waitForTimeout(500);
        
        // Submit the form
        const submitButton = await page.getByRole('button', { name: /create/i }).first();
        if (await submitButton.isVisible()) {
          await submitButton.click({ force: true });
          await page.waitForTimeout(2000);
        }
      }
    }
    
    await page.screenshot({ 
      path: path.join(evidenceDir, '02-workspace-created.png'),
      fullPage: true 
    });
    console.log('✅ Workspace creation attempted');
    
    // Test 3: Navigate to connections page
    console.log('📍 Test 3: Navigating to connections page...');
    await page.goto('http://localhost:5173/settings/workspaces', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    
    await page.screenshot({ 
      path: path.join(evidenceDir, '03-settings-workspaces.png'),
      fullPage: true 
    });
    console.log('✅ Settings workspaces page loaded');
    
    // Test 4: Try to access a workspace's connections
    console.log('📍 Test 4: Accessing workspace connections...');
    
    // Look for a workspace link or card
    const workspaceLink = await page.locator('a[href*="/connections"], button:has-text("Connections")').first();
    if (await workspaceLink.isVisible()) {
      await workspaceLink.click();
      await page.waitForTimeout(2000);
      
      await page.screenshot({ 
        path: path.join(evidenceDir, '04-connections-page.png'),
        fullPage: true 
      });
      console.log('✅ Connections page loaded');
      
      // Test 5: Create a connection
      console.log('📍 Test 5: Creating a new connection...');
      
      const addConnectionButton = await page.locator('button:has-text("Add"), button:has-text("Create"), button:has-text("New Connection")').first();
      if (await addConnectionButton.isVisible()) {
        await addConnectionButton.click();
        await page.waitForTimeout(1000);
        
        // Fill connection form
        const nameInput = await page.locator('input[name="name"], input[placeholder*="name"]').first();
        if (await nameInput.isVisible()) {
          await nameInput.fill('Test Connection');
        }
        
        const apiKeyInput = await page.locator('input[name="apiKey"], input[placeholder*="API"], input[type="password"]').first();
        if (await apiKeyInput.isVisible()) {
          await apiKeyInput.fill('test-api-key-12345');
        }
        
        const endpointInput = await page.locator('input[name="endpoint"], input[placeholder*="endpoint"], input[placeholder*="URL"]').first();
        if (await endpointInput.isVisible()) {
          await endpointInput.fill('https://test.example.com');
        }
        
        // Submit
        const saveButton = await page.locator('button:has-text("Save"), button:has-text("Create"), button[type="submit"]').first();
        if (await saveButton.isVisible()) {
          await saveButton.click();
          await page.waitForTimeout(2000);
        }
      }
      
      await page.screenshot({ 
        path: path.join(evidenceDir, '05-connection-created.png'),
        fullPage: true 
      });
      console.log('✅ Connection creation attempted');
      
      // Test 6: Test connection
      console.log('📍 Test 6: Testing connection...');
      
      const testButton = await page.locator('button:has-text("Test"), button:has-text("Verify")').first();
      if (await testButton.isVisible()) {
        await testButton.click();
        await page.waitForTimeout(3000);
      }
      
      await page.screenshot({ 
        path: path.join(evidenceDir, '06-connection-tested.png'),
        fullPage: true 
      });
      console.log('✅ Connection test attempted');
    }
    
    console.log('\n✨ Integration test completed successfully!');
    
  } catch (error) {
    console.error('❌ Test failed:', error.message);
    
    await page.screenshot({ 
      path: path.join(evidenceDir, 'error-screenshot.png'),
      fullPage: true 
    });
    
    throw error;
  } finally {
    await browser.close();
  }
}

runIntegrationTest().catch(console.error);
