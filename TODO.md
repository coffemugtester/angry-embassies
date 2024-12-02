# TODO:

[] add pictures to mongo

[] Add InitHandler and isolate commands
[] Create runService method

[x] Add a server

[x] Add read method to the repository

[x] Make cmd directory; should contain 2 modules api-server and cli

[x] Populate all Spanish embassies properly (tweak the input process)

[x] Extract initialization of dependencies out of the CLI setup and initialize the app properly

[x] install raspberry pi os on the sd card
https://www.amazon.de/-/en/Micro-Reader-Memory-External-RS-MMC-gray/dp/B087QG75L7/ref=sr_1_1?crid=ZIVFZD8J3K0K&dib=eyJ2IjoiMSJ9.fRdmcvVUMsaRIuRv_VsJE8dAEL03mbsBy0Gy8Djq1DY6-63rz84YcZn43CDur4iVthUePRzsumYJ2Ymty7rA3JVg-KThSmAl5vCI9uhoEl16hvMwzOOg2DlKShFH_XvNBSe3kBwx-FAXcf5zWIEGIfB1sVHxwRE3jG_3fcvv7iN-zGExME8Heg-veYUJrU1LUgRXSAXWVroszUCSh3e7xGtwqA37eiaHgYN2r8OWclY.7EH6gPLfeNiXeIkv0QzYoJUtIH8dyP0K8YGtb18J7Zc&dib_tag=se&keywords=usb+sdcard&qid=1724681407&sprefix=usb+sdcar%2Caps%2C108&sr=8-1

[x] 1 ssh access to the raspberry pi

Automate CI/CD pipeline for the project (avoid using GitHub actions for now):

    [x] install go on the raspberry pi
    [x] install docker on the raspberry pi
    [x] run the service in a docker container on the raspberry pi
    [x] install jenkins and enabble it
    [x] write cron job script to check for changes in the repository
    [x] enable cron job
    write a script/use jenkins to run tests and rebuild the docker image and launch the service

    pull the code *securely* from the repository
Here are ways to implement secure practices in your cron job:

1. Use SSH Keys (instead of HTTP):
   Set up SSH keys for authentication, as SSH is more secure than HTTP because it encrypts the connection. Make sure your private key is stored securely and without a passphrase if automation is required.
   In your cron job script, use git pull over SSH (e.g., git@github.com:user/repo.git) to leverage the security of SSH keys.
2. Restrict Permissions on SSH Keys:
   Ensure the SSH key you use has the minimum permissions necessary for pulling (not pushing) code, ideally with read-only access.
   Limit file permissions for the SSH key file itself to prevent unauthorized access:
   bash
   Code kopieren
   chmod 600 ~/.ssh/id_rsa
3. Environment Variables for Sensitive Data:
   Avoid hardcoding sensitive information (like tokens or keys) directly in the cron script. Use environment variables if needed and load them securely within the script.
4. Pull Over VPN or Firewall-Restricted Access (Optional):
   If applicable, restrict the server’s outbound access to Git only over a VPN or through a firewall rule that limits access to the specific IP range of the Git provider.
5. Verify Repository Authenticity:
   Use Git’s gpg verification (if you can configure it) to ensure the code comes from a trusted source.
   Example Cron Job with SSH Pull:
   bash
   Code kopieren
   */5 * * * * /usr/bin/git -C /path/to/repo pull origin main > /path/to/logfile.log 2>&1
   These measures can help ensure that your cron job script pulls code securely without risking unauthorized access.

https://www.amazon.de/-/en/SanDisk-microSDXC-Smartphones-Transmission-RescuePRO/dp/B09X7BK27V/ref=sr_1_3?crid=1GW0YGXI8WL95&dib=eyJ2IjoiMSJ9.hNNT_coZ6zMdsZbsxVtWxb7GZ4rhT_Hqqnf_jkMGcAJUB9xbaK79CMweIVvbuwDOqSDdB38KgWS53f_r2JOI5m2R6iXbd97tHhZQpCrDXY1on77KrNSmzWhaZKNPrqbVphRHWXDCTQP0EsvA2PVIRAWV_81IP5jDMuL-UaYAK0v7S3Wo32ZF-TANQZ3tvh2yA9Gt8whnZnnyHLY-5ZRPFPLZrITLNgaL7itCDSgx9VE.pePYmTE7fmGj8wW5s6i3q0ws9MLN8rK3uBE4PVsb084&dib_tag=se&keywords=sdcard&qid=1724681471&sprefix=sdcard+%2Caps%2C110&sr=8-3

https://www.amazon.de/-/en/Pimoroni-NVMe-Base-2230-Raspberry-multi-coloured/dp/B0CTK1RLN5/ref=sr_1_4?crid=1GNGFWSDWR9FP&dib=eyJ2IjoiMSJ9.j1zgJqQiQixbpxsQcCbwOg90lMVg3TACcifX0JFisoc7XHHiVYTxcrsvIuLvxOfU_4mRb_aERrGNkb61gvrORBs611hBtkXDn4U5fE3oeRTtOdJcGUpvbKeen3FNw0wJ1gomoUCf2b9PuyG5WX3BHD0MUrLYccdF0UAbTHbUgm5QMk6IDsQJLCwk9skpCULAmAPk428Xp8HVrfyOv1Vg7bQzbF-jsQ6OX_gZvui3dVQ.Ip_vHrF-iQbC_aYgUzu46K7Oy3fQFtFUtK2dS8e2b0g&dib_tag=se&keywords=raspberry%2Bpi%2B5%2Bexternal%2Bdisk&qid=1724681587&sprefix=raspberry%2Bpi%2B5%2Bexternal%2Bdisk%2Caps%2C106&sr=8-4&th=1

https://www.amazon.de/-/en/Crucial-Internal-Compatible-Laptop-Desktop/dp/B0BYWB6237/ref=sr_1_3?crid=2MXKJNSUDN3JR&dib=eyJ2IjoiMSJ9.vYOYfnTlcKt_hLX4i8nA9d_60-eGkaBbysd-U3BiA1YnQxpI_E6SptGiclvRk9EviEzaOYHPh6NpOQ6rcDrksXzplnUR14WFA5iQgXFRhUogg6wtr-L-35J0S_0XbCQKs-eXZL-r6lC2n9Fh5ATKAMbFit7-iO5BI3Wdy7j-AHe5ylXgr-vEGd-M686DFCrCUOdeqIZLd8ZNkFu6eg2CHmXAnPC2jjIwRdxv2pR-WjU.gaNxDBX36W-BxgJfCZIkGoqNxNKLObxUgkzrEbXHC0s&dib_tag=se&keywords=m.2+2280&qid=1724681703&sprefix=m.2+2280%2Caps%2C110&sr=8-3
