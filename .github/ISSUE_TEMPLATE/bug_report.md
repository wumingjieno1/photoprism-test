---
name: Bug Report
about: Report a new and clearly identified bug that must be fixed directly in the application
title: 'Category: Clear and concise title (PLEASE CHANGE)'
labels: bug
assignees: ''

---

STOP! BEFORE YOU PROCEED: ARE YOU 100% SURE THAT THIS IS NOT A TECHNICAL SUPPORT REQUEST AND/OR A PROBLEM WITH ANOTHER SOFTWARE YOU ARE USING, SUCH AS NGINX?

- We kindly ask you not to report bugs via GitHub Issues **unless you are certain to have found a fully reproducible and previously unreported issue** that must be fixed directly in the app
- When reporting a problem, **always include the software versions you are using and other information about your environment** such as browser, browser plugins, operating system, storage type, memory size, and processor
- Note that all issue **subscribers receive an email notification** from GitHub for each new comment, so these should only be used for sharing important information and not for personal discussions/questions
- Contact us or a community member if you need help, it could be a local configuration problem, or a misunderstanding in how the software works
- This gives our team the opportunity to improve the docs and provide best-in-class support to you, instead of handling unclear/duplicate bug reports or triggering a flood of notifications by responding to comments

USE THE TROUBLESHOOTING CHECKLISTS IN OUR DOCS TO QUICKLY DIAGNOSE AND FIX COMMON PROBLEMS:

  <https://docs.photoprism.app/getting-started/troubleshooting/>

BEFORE OPENING AN ISSUE, MAKE SURE TO SEARCH EXISTING BUG REPORTS AND FEATURE REQUESTS:

  <https://github.com/photoprism/photoprism/issues>

FOR GENERAL QUESTIONS, TECHNICAL SUPPORT, AND TO GET TO KNOW OTHER COMMUNITY MEMBERS:

  <https://github.com/photoprism/photoprism/discussions>

PLEASE PROCEED ONLY IF YOU ARE 100% SURE THAT THIS IS NOT A TECHNICAL SUPPORT REQUEST AND/OR A PROBLEM WITH ANOTHER SOFTWARE YOU ARE USING.

**THANK YOU FOR YOUR CAREFUL CONSIDERATION!**

===============================================================================

**What does not work as described in the documentation?**

"No photos found" is not a sufficient description. Please be more specific and explain which part of the software has a bug and needs to be fixed. Do not report known issues or features not yet implemented as bugs.

**How can we reproduce it?**

Steps to reproduce the behavior:

1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

If reporting an Import, Indexing, or Performance issue, please include the number and type of pictures in your library,
as well as any configuration options you have changed e.g. for thumbnail quality.

**What behavior do you expect?**

A clear and concise description of what you expected to happen.

**What could be the cause of your problem?**

Try to determine the cause of your problem before submitting a bug report:
https://docs.photoprism.app/getting-started/troubleshooting/

**Can you provide us with example files for testing, error logs, or screenshots?**

Please add example files or screenshots that help to reproduce your problem.
You may also send files via email or share a download link:
https://photoprism.app/contact

Learn how to diagnose frontend issues:
https://docs.photoprism.app/getting-started/troubleshooting/browsers/

NOTE:
- You have to provide sample files and logs if it is an IMPORT, INDEXING, or METADATA issue, otherwise we will not be able to process your report
- If it is an IMPORT issue, please also provide an archive with affected files before importing them so that it's possible to reproduce your issue

**Which software versions do you use?**

- PhotoPrism Architecture & Build Number (AMD64, ARM64, ARMv7,...)
- Database Type & Version (MariaDB, MySQL, SQLite,...)
- Operating System Types & Versions (Linux, Windows, Android,...)
- Browser Types & Versions (Firefox, Chrome, Safari on iPhone,...)
- Browser Plugins & Ad Blockers (if any)

The app version / build number can be found in *Settings* when you scroll to the bottom.
MySQL 8 is not officially supported anymore, see System Requirements.

NOTE:
- Always provide database and operating system details if it is a Backend, Import, or Indexing issue
- Always provide web browser and operating system details if it is a Frontend issue
- If it is a Performance problem, you must provide ALL information

**On what kind of device is PhotoPrism installed?**

This is especially important if you are reporting a Performance, Import, or Indexing problem. You can skip this if you are reporting a problem found on our public demo, or if it is a completely unrelated issue, such as a broken page layout.

- Device / Processor Type (Raspberry Pi 4, Intel Core i7-3770, AMD Ryzen 7 3800X,...)
- Physical Memory & Swap Space (in GB)
- Storage Type (HDD, SSD, RAID, USB, Network Storage,...) 
- anything else that might be helpful

NOTE:
- Always provide device, memory, and storage details when you have a Backend, Performance, Import, or Indexing issue

**Do you use a Reverse Proxy, Firewall, VPN, or CDN?**

If yes, please specify type and version. You can skip this if you are reporting a completely unrelated issue.

NOTE:
- Always provide this information when you have a Reliability, Performance, or Frontend problem, such as failed uploads, connection errors, broken thumbnails, or video playback issues
- If you are using NGINX, also provide its configuration and/or consider asking the NGINX community for advice as we do not specialize in supporting their product: https://docs.photoprism.app/getting-started/proxies/nginx/
