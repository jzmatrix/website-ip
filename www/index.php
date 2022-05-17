<?php
	$ip = $_SERVER{'REMOTE_ADDR'};
	if (isset($_SERVER['HTTP_X_FORWARDED_FOR']))
	{
		$forwardedFor = $_SERVER['HTTP_X_FORWARDED_FOR'];
		$ipList = explode('/,/', $forwardedFor);
		$myIP = $ipList[0];
		print $myIP;
	}
	else
	{
		print $ip;
	}
	
	#print "$ip :: $forwardedFor :: $myIP";
	
	# phpinfo();
?>
