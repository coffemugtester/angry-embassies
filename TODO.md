# TODO:
test change

[x] install raspberry pi os on the sd card
https://www.amazon.de/-/en/Micro-Reader-Memory-External-RS-MMC-gray/dp/B087QG75L7/ref=sr_1_1?crid=ZIVFZD8J3K0K&dib=eyJ2IjoiMSJ9.fRdmcvVUMsaRIuRv_VsJE8dAEL03mbsBy0Gy8Djq1DY6-63rz84YcZn43CDur4iVthUePRzsumYJ2Ymty7rA3JVg-KThSmAl5vCI9uhoEl16hvMwzOOg2DlKShFH_XvNBSe3kBwx-FAXcf5zWIEGIfB1sVHxwRE3jG_3fcvv7iN-zGExME8Heg-veYUJrU1LUgRXSAXWVroszUCSh3e7xGtwqA37eiaHgYN2r8OWclY.7EH6gPLfeNiXeIkv0QzYoJUtIH8dyP0K8YGtb18J7Zc&dib_tag=se&keywords=usb+sdcard&qid=1724681407&sprefix=usb+sdcar%2Caps%2C108&sr=8-1

[x] 1 ssh access to the raspberry pi

Automate CI/CD pipeline for the project (avoid using GitHub actions for now):

    [x] install go on the raspberry pi
    [x] install docker on the raspberry pi
    [x] run the service in a docker container on the raspberry pi
    [x] install jenkins and enabble it
    [x] write cron job script to check for changes in the repository
    enable cron job
    write a script/use jenkins to run tests and rebuild the docker image and launch the service

    pull the code *securely* from the repository
    check diffs
    reload the service (rebuild docker image and launch it)
    (put everything on scripts and store them)
    (use bash scripts and keep them versioned and close to the project)

https://www.amazon.de/-/en/SanDisk-microSDXC-Smartphones-Transmission-RescuePRO/dp/B09X7BK27V/ref=sr_1_3?crid=1GW0YGXI8WL95&dib=eyJ2IjoiMSJ9.hNNT_coZ6zMdsZbsxVtWxb7GZ4rhT_Hqqnf_jkMGcAJUB9xbaK79CMweIVvbuwDOqSDdB38KgWS53f_r2JOI5m2R6iXbd97tHhZQpCrDXY1on77KrNSmzWhaZKNPrqbVphRHWXDCTQP0EsvA2PVIRAWV_81IP5jDMuL-UaYAK0v7S3Wo32ZF-TANQZ3tvh2yA9Gt8whnZnnyHLY-5ZRPFPLZrITLNgaL7itCDSgx9VE.pePYmTE7fmGj8wW5s6i3q0ws9MLN8rK3uBE4PVsb084&dib_tag=se&keywords=sdcard&qid=1724681471&sprefix=sdcard+%2Caps%2C110&sr=8-3

https://www.amazon.de/-/en/Pimoroni-NVMe-Base-2230-Raspberry-multi-coloured/dp/B0CTK1RLN5/ref=sr_1_4?crid=1GNGFWSDWR9FP&dib=eyJ2IjoiMSJ9.j1zgJqQiQixbpxsQcCbwOg90lMVg3TACcifX0JFisoc7XHHiVYTxcrsvIuLvxOfU_4mRb_aERrGNkb61gvrORBs611hBtkXDn4U5fE3oeRTtOdJcGUpvbKeen3FNw0wJ1gomoUCf2b9PuyG5WX3BHD0MUrLYccdF0UAbTHbUgm5QMk6IDsQJLCwk9skpCULAmAPk428Xp8HVrfyOv1Vg7bQzbF-jsQ6OX_gZvui3dVQ.Ip_vHrF-iQbC_aYgUzu46K7Oy3fQFtFUtK2dS8e2b0g&dib_tag=se&keywords=raspberry%2Bpi%2B5%2Bexternal%2Bdisk&qid=1724681587&sprefix=raspberry%2Bpi%2B5%2Bexternal%2Bdisk%2Caps%2C106&sr=8-4&th=1

https://www.amazon.de/-/en/Crucial-Internal-Compatible-Laptop-Desktop/dp/B0BYWB6237/ref=sr_1_3?crid=2MXKJNSUDN3JR&dib=eyJ2IjoiMSJ9.vYOYfnTlcKt_hLX4i8nA9d_60-eGkaBbysd-U3BiA1YnQxpI_E6SptGiclvRk9EviEzaOYHPh6NpOQ6rcDrksXzplnUR14WFA5iQgXFRhUogg6wtr-L-35J0S_0XbCQKs-eXZL-r6lC2n9Fh5ATKAMbFit7-iO5BI3Wdy7j-AHe5ylXgr-vEGd-M686DFCrCUOdeqIZLd8ZNkFu6eg2CHmXAnPC2jjIwRdxv2pR-WjU.gaNxDBX36W-BxgJfCZIkGoqNxNKLObxUgkzrEbXHC0s&dib_tag=se&keywords=m.2+2280&qid=1724681703&sprefix=m.2+2280%2Caps%2C110&sr=8-3
