package demo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strings"
	"time"
)

const (
	emailItAPIURL = "https://api.emailit.com/v1/emails"
	maxRetries    = 3
	baseDelay     = 1 * time.Second
)

// EmailService handles sending emails via EmailIt
type EmailService struct {
	config *Config
	client *http.Client
}

// NewEmailService creates a new email service instance
func NewEmailService(config *Config) *EmailService {
	return &EmailService{
		config: config,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SendDemoEmails sends both the thank you and notification emails
func (es *EmailService) SendDemoEmails(req *DemoRequest, ipAddress string) error {
	// Send thank you email to prospect
	if err := es.sendThankYouEmail(req); err != nil {
		return fmt.Errorf("failed to send thank you email: %w", err)
	}
	
	// Send notification email to admin
	if err := es.sendNotificationEmail(req, ipAddress); err != nil {
		return fmt.Errorf("failed to send notification email: %w", err)
	}
	
	return nil
}

// sendThankYouEmail sends a confirmation email to the prospect
func (es *EmailService) sendThankYouEmail(req *DemoRequest) error {
	subject := "Your Blue Enterprise Demo Request Received"
	
	htmlTemplate := `
<!doctype html>
<html lang="en">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>Your Blue Enterprise Demo Request</title>
    <style media="all" type="text/css">
    @media only screen and (max-width: 640px) {
        .main p,
        .main td,
        .main span {
            font-size: 16px !important;
        }
        .wrapper {
            padding: 8px !important;
        }
        .content {
            padding: 0 !important;
        }
        .container {
            padding: 0 !important;
            padding-top: 8px !important;
            width: 100% !important;
        }
        .main {
            border-left-width: 0 !important;
            border-radius: 0 !important;
            border-right-width: 0 !important;
        }
        .btn table {
            max-width: 100% !important;
            width: 100% !important;
        }
        .btn a {
            font-size: 16px !important;
            max-width: 100% !important;
            width: 100% !important;
        }
    }
    @media all {
        .ExternalClass {
            width: 100%;
        }
        .ExternalClass,
        .ExternalClass p,
        .ExternalClass span,
        .ExternalClass font,
        .ExternalClass td,
        .ExternalClass div {
            line-height: 100%;
        }
        .apple-link a {
            color: inherit !important;
            font-family: inherit !important;
            font-size: inherit !important;
            font-weight: inherit !important;
            line-height: inherit !important;
            text-decoration: none !important;
        }
        #MessageViewBody a {
            color: inherit;
            text-decoration: none;
            font-size: inherit;
            font-family: inherit;
            font-weight: inherit;
            line-height: inherit;
        }
    }
    </style>
</head>
<body style="font-family: 'Helvetica Neue', Helvetica, sans-serif; -webkit-font-smoothing: antialiased; font-size: 16px; line-height: 1.3; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; background-color: #f4f5f6; margin: 0; padding: 0;">
    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f4f5f6; width: 100%;" width="100%" bgcolor="#f4f5f6">
        <tr>
            <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">&nbsp;</td>
            <td class="container" style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; max-width: 600px; padding: 0; padding-top: 24px; width: 600px; margin: 0 auto;" width="600" valign="top">
                <div class="content" style="box-sizing: border-box; display: block; margin: 0 auto; max-width: 600px; padding: 0;">
                    
                    <!-- START CENTERED WHITE CONTAINER -->
                    <span class="preheader" style="color: transparent; display: none; height: 0; max-height: 0; max-width: 0; opacity: 0; overflow: hidden; mso-hide: all; visibility: hidden; width: 0;">Thank you for requesting a demo of Blue Enterprise. We'll contact you within 24 hours.</span>
                    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="main" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background: #ffffff; border: 1px solid #eaebed; border-radius: 16px; width: 100%;" width="100%">
                        
                        <!-- START MAIN CONTENT AREA -->
                        <tr>
                            <td class="wrapper" style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; box-sizing: border-box; padding: 24px;" valign="top">
                                
                                <!-- Logo/Header -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; margin-bottom: 24px;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; text-align: center;" valign="top" align="center">
                                            <h1 style="color: #00a0d2; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 16px; font-size: 24px;">Thank You for Your Interest in Blue!</h1>
                                        </td>
                                    </tr>
                                </table>
                                
                                <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; font-weight: normal; margin: 0; margin-bottom: 16px;">Hi {{.FullName}},</p>
                                
                                <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; font-weight: normal; margin: 0; margin-bottom: 16px;">Thank you for requesting a demo of Blue Enterprise. We've received your request and our team is reviewing your requirements.</p>
                                
                                <!-- Details Box -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; margin-bottom: 24px; margin-top: 24px;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; background-color: #f8f9fa; border-radius: 8px; padding: 16px;" valign="top" bgcolor="#f8f9fa">
                                            <h3 style="color: #06090f; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 12px; font-size: 18px;">Your Demo Request Details:</h3>
                                            <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Company:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.Company}}</td>
                                                </tr>
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Team Size:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.CompanySize}} employees</td>
                                                </tr>
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Use Case:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.UseCase}}</td>
                                                </tr>
                                            </table>
                                        </td>
                                    </tr>
                                </table>
                                
                                <h3 style="color: #06090f; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 12px; margin-top: 24px; font-size: 18px;">What Happens Next:</h3>
                                
                                <ol style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; font-weight: normal; margin: 0; margin-bottom: 16px; padding-left: 24px;">
                                    <li style="margin-bottom: 8px;">Our enterprise team will review your specific needs</li>
                                    <li style="margin-bottom: 8px;">We'll reach out within 24 hours to schedule your personalized demo</li>
                                    <li style="margin-bottom: 8px;">The demo will focus on your exact use cases and workflows</li>
                                </ol>
                                
                                <!-- CTA Button -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; box-sizing: border-box; width: 100%; min-width: 100%; margin-top: 24px; margin-bottom: 24px;" width="100%">
                                    <tbody>
                                        <tr>
                                            <td align="center" style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; padding-bottom: 16px;" valign="top">
                                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: auto;">
                                                    <tbody>
                                                        <tr>
                                                            <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; border-radius: 8px; text-align: center; background-color: #00a0d2;" valign="top" align="center" bgcolor="#00a0d2">
                                                                <a href="https://blue.cc/docs" target="_blank" style="border: solid 2px #00a0d2; border-radius: 8px; box-sizing: border-box; cursor: pointer; display: inline-block; font-size: 16px; font-weight: bold; margin: 0; padding: 12px 24px; text-decoration: none; text-transform: capitalize; background-color: #00a0d2; border-color: #00a0d2; color: #ffffff;">Explore Documentation</a>
                                                            </td>
                                                        </tr>
                                                    </tbody>
                                                </table>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                                
                                <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 16px; color: #6c757d;">Feel free to explore our documentation while you wait to learn more about Blue's capabilities.</p>
                                
                                <hr style="border: 0; border-top: 1px solid #eaebed; margin: 24px 0;">
                                
                                <!-- Footer -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #9a9ea6; text-align: center; padding-top: 16px;" valign="top" align="center">
                                            <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 4px; color: #9a9ea6;">Best regards,<br>The Blue Team</p>
                                            <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 12px; font-weight: normal; margin: 0; margin-bottom: 4px; color: #9a9ea6;">
                                                Blue - Enterprise Process Automation<br>
                                                <a href="https://blue.cc" style="color: #00a0d2; text-decoration: underline;">blue.cc</a> | <a href="mailto:enterprise@blue.cc" style="color: #00a0d2; text-decoration: underline;">enterprise@blue.cc</a>
                                            </p>
                                        </td>
                                    </tr>
                                </table>
                                
                            </td>
                        </tr>
                        
                    </table>
                    
                </div>
            </td>
            <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">&nbsp;</td>
        </tr>
    </table>
</body>
</html>`
	
	textTemplate := `Hi {{.FullName}},

Thank you for requesting a demo of Blue Enterprise. We've received your request and our team is reviewing your requirements.

Your Demo Request Details:
• Company: {{.Company}}
• Team Size: {{.CompanySize}}
• Primary Use Case: {{.UseCase}}

What Happens Next:
1. Our enterprise team will review your specific needs
2. We'll reach out within 24 hours to schedule your personalized demo
3. The demo will focus on your exact use cases and workflows

In the meantime, feel free to explore our documentation at blue.cc/docs to learn more about Blue's capabilities.

Best regards,
The Blue Team

---
Blue - Enterprise Process Automation
blue.cc | enterprise@blue.cc`
	
	// Parse and execute templates
	html, err := executeTemplate(htmlTemplate, req)
	if err != nil {
		return err
	}
	
	text, err := executeTemplate(textTemplate, req)
	if err != nil {
		return err
	}
	
	payload := EmailPayload{
		From:    fmt.Sprintf("%s <%s>", es.config.EmailItFromName, es.config.EmailItFromEmail),
		To:      req.Email,
		Subject: subject,
		HTML:    html,
		Text:    text,
	}
	
	return es.sendWithRetry(payload)
}

// sendNotificationEmail sends a notification to the admin
func (es *EmailService) sendNotificationEmail(req *DemoRequest, ipAddress string) error {
	subject := fmt.Sprintf("New Demo Request: %s - %s", req.Company, req.FullName)
	
	// Format use case for display
	useCase := strings.ReplaceAll(req.UseCase, "-", " ")
	useCase = strings.Title(useCase)
	
	htmlTemplate := `
<!doctype html>
<html lang="en">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>New Demo Request: {{.Company}}</title>
    <style media="all" type="text/css">
    @media only screen and (max-width: 640px) {
        .main p,
        .main td,
        .main span {
            font-size: 16px !important;
        }
        .wrapper {
            padding: 8px !important;
        }
        .content {
            padding: 0 !important;
        }
        .container {
            padding: 0 !important;
            padding-top: 8px !important;
            width: 100% !important;
        }
        .main {
            border-left-width: 0 !important;
            border-radius: 0 !important;
            border-right-width: 0 !important;
        }
    }
    @media all {
        .ExternalClass {
            width: 100%;
        }
        .ExternalClass,
        .ExternalClass p,
        .ExternalClass span,
        .ExternalClass font,
        .ExternalClass td,
        .ExternalClass div {
            line-height: 100%;
        }
        .apple-link a {
            color: inherit !important;
            font-family: inherit !important;
            font-size: inherit !important;
            font-weight: inherit !important;
            line-height: inherit !important;
            text-decoration: none !important;
        }
        #MessageViewBody a {
            color: inherit;
            text-decoration: none;
            font-size: inherit;
            font-family: inherit;
            font-weight: inherit;
            line-height: inherit;
        }
    }
    </style>
</head>
<body style="font-family: 'Helvetica Neue', Helvetica, sans-serif; -webkit-font-smoothing: antialiased; font-size: 16px; line-height: 1.3; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; background-color: #f4f5f6; margin: 0; padding: 0;">
    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f4f5f6; width: 100%;" width="100%" bgcolor="#f4f5f6">
        <tr>
            <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">&nbsp;</td>
            <td class="container" style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; max-width: 600px; padding: 0; padding-top: 24px; width: 600px; margin: 0 auto;" width="600" valign="top">
                <div class="content" style="box-sizing: border-box; display: block; margin: 0 auto; max-width: 600px; padding: 0;">
                    
                    <!-- START CENTERED WHITE CONTAINER -->
                    <span class="preheader" style="color: transparent; display: none; height: 0; max-height: 0; max-width: 0; opacity: 0; overflow: hidden; mso-hide: all; visibility: hidden; width: 0;">New demo request from {{.Company}} - {{.FullName}}</span>
                    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="main" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background: #ffffff; border: 1px solid #eaebed; border-radius: 16px; width: 100%;" width="100%">
                        
                        <!-- START MAIN CONTENT AREA -->
                        <tr>
                            <td class="wrapper" style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; box-sizing: border-box; padding: 24px;" valign="top">
                                
                                <!-- Header with urgency indicator -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; margin-bottom: 24px;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">
                                            <span style="display: inline-block; padding: 4px 8px; background-color: #28a745; color: white; border-radius: 4px; font-size: 12px; font-weight: bold; text-transform: uppercase; margin-bottom: 8px;">New Lead</span>
                                            <h1 style="color: #00a0d2; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 8px; font-size: 24px;">Enterprise Demo Request</h1>
                                            <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 18px; font-weight: 600; margin: 0; color: #333;">{{.Company}}</p>
                                        </td>
                                    </tr>
                                </table>
                                
                                <!-- Contact Information -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; margin-bottom: 24px;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; background-color: #f8f9fa; border-radius: 8px; padding: 16px;" valign="top" bgcolor="#f8f9fa">
                                            <h3 style="color: #06090f; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 12px; font-size: 16px;">Contact Information</h3>
                                            <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0; width: 30%;" valign="top">Name:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.FullName}}</td>
                                                </tr>
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Email:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">
                                                        <a href="mailto:{{.Email}}" style="color: #00a0d2; text-decoration: underline;">{{.Email}}</a>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Job Title:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.JobTitle}}</td>
                                                </tr>
                                                {{if .Phone}}
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Phone:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.Phone}}</td>
                                                </tr>
                                                {{end}}
                                            </table>
                                        </td>
                                    </tr>
                                </table>
                                
                                <!-- Company Details -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; margin-bottom: 24px;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; background-color: #fff4e5; border-radius: 8px; padding: 16px;" valign="top" bgcolor="#fff4e5">
                                            <h3 style="color: #06090f; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 12px; font-size: 16px;">Company Details</h3>
                                            <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0; width: 30%;" valign="top">Company:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.Company}}</td>
                                                </tr>
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Size:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.CompanySize}} employees</td>
                                                </tr>
                                                <tr>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; color: #6c757d; padding: 4px 0;" valign="top">Use Case:</td>
                                                    <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; vertical-align: top; font-weight: 600; padding: 4px 0;" valign="top">{{.UseCaseFormatted}}</td>
                                                </tr>
                                            </table>
                                        </td>
                                    </tr>
                                </table>
                                
                                {{if .Message}}
                                <!-- Message from Prospect -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; margin-bottom: 24px;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; background-color: #e3f2fd; border-radius: 8px; padding: 16px;" valign="top" bgcolor="#e3f2fd">
                                            <h3 style="color: #06090f; font-family: 'Helvetica Neue', Helvetica, sans-serif; font-weight: 600; line-height: 1.4; margin: 0; margin-bottom: 12px; font-size: 16px;">Message from Prospect</h3>
                                            <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 14px; font-weight: normal; margin: 0; color: #333; white-space: pre-wrap;">{{.Message}}</p>
                                        </td>
                                    </tr>
                                </table>
                                {{end}}
                                
                                <!-- Action Button -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; box-sizing: border-box; width: 100%; min-width: 100%; margin-bottom: 24px;" width="100%">
                                    <tbody>
                                        <tr>
                                            <td align="center" style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">
                                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: auto;">
                                                    <tbody>
                                                        <tr>
                                                            <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top; border-radius: 8px; text-align: center; background-color: #00a0d2;" valign="top" align="center" bgcolor="#00a0d2">
                                                                <a href="mailto:{{.Email}}" style="border: solid 2px #00a0d2; border-radius: 8px; box-sizing: border-box; cursor: pointer; display: inline-block; font-size: 16px; font-weight: bold; margin: 0; padding: 12px 24px; text-decoration: none; text-transform: capitalize; background-color: #00a0d2; border-color: #00a0d2; color: #ffffff;">Reply to {{.FullName}}</a>
                                                            </td>
                                                        </tr>
                                                    </tbody>
                                                </table>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                                
                                <hr style="border: 0; border-top: 1px solid #eaebed; margin: 24px 0;">
                                
                                <!-- Meta Information -->
                                <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                                    <tr>
                                        <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 12px; vertical-align: top; color: #9a9ea6;" valign="top">
                                            <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 12px; font-weight: normal; margin: 0; margin-bottom: 4px; color: #9a9ea6;">
                                                <strong>Submitted:</strong> {{.Timestamp}}<br>
                                                <strong>IP Address:</strong> {{.IPAddress}}<br>
                                                <strong>Form:</strong> Enterprise Demo Request
                                            </p>
                                            <p style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 12px; font-weight: normal; margin: 0; margin-top: 12px; color: #9a9ea6; font-style: italic;">
                                                💡 Tip: Reply directly to this email to respond to {{.FullName}}. Their email is set as the reply-to address.
                                            </p>
                                        </td>
                                    </tr>
                                </table>
                                
                            </td>
                        </tr>
                        
                    </table>
                    
                </div>
            </td>
            <td style="font-family: 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">&nbsp;</td>
        </tr>
    </table>
</body>
</html>`
	
	textTemplate := `New enterprise demo request received:

Contact Information:
• Name: {{.FullName}}
• Email: {{.Email}}
• Company: {{.Company}}
• Job Title: {{.JobTitle}}
{{if .Phone}}• Phone: {{.Phone}}{{end}}

Company Details:
• Size: {{.CompanySize}} employees
• Primary Use Case: {{.UseCaseFormatted}}

{{if .Message}}Message from Prospect:
{{.Message}}{{end}}

---
Submitted: {{.Timestamp}}
IP Address: {{.IPAddress}}

Reply directly to this email to respond to {{.FullName}}.`
	
	// Create data with additional fields
	data := struct {
		*DemoRequest
		UseCaseFormatted string
		Timestamp        string
		IPAddress        string
	}{
		DemoRequest:      req,
		UseCaseFormatted: useCase,
		Timestamp:        time.Now().Format("2006-01-02 15:04:05 MST"),
		IPAddress:        ipAddress,
	}
	
	// Parse and execute templates
	html, err := executeTemplate(htmlTemplate, data)
	if err != nil {
		return err
	}
	
	text, err := executeTemplate(textTemplate, data)
	if err != nil {
		return err
	}
	
	payload := EmailPayload{
		From:    fmt.Sprintf("%s <%s>", es.config.EmailItFromName, es.config.EmailItFromEmail),
		To:      es.config.NotificationEmail,
		ReplyTo: req.Email,
		Subject: subject,
		HTML:    html,
		Text:    text,
	}
	
	return es.sendWithRetry(payload)
}

// sendWithRetry sends an email with exponential backoff retry logic
func (es *EmailService) sendWithRetry(payload EmailPayload) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		err := es.sendEmail(payload)
		if err == nil {
			return nil
		}
		
		lastErr = err
		
		// Don't retry on the last attempt
		if i < maxRetries-1 {
			delay := baseDelay * time.Duration(math.Pow(2, float64(i)))
			time.Sleep(delay)
		}
	}
	
	return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

// sendEmail sends a single email via EmailIt API
func (es *EmailService) sendEmail(payload EmailPayload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	
	req, err := http.NewRequest("POST", emailItAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	req.Header.Set("Authorization", "Bearer "+es.config.EmailItAPIKey)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := es.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("EmailIt API returned status %d", resp.StatusCode)
	}
	
	return nil
}

// executeTemplate executes a template with the given data
func executeTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("email").Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}
	
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	
	return buf.String(), nil
}