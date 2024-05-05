package com.example.smoking.ui.theme.page

import android.app.Activity
import android.os.Bundle
import android.widget.Button
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.safeContentPadding
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.RectangleShape
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.smoking.R

@Composable
fun PageScreen(navController: NavHostController) {
    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally

    ) {
        Button(onClick = { navController.navigate("textPage") }) {
            Text("Open Text Page")
        }
    }
}
@Preview
@Composable
fun PageScreenPreview() {
    val navController = rememberNavController()
    PageScreen(navController)
}

@Preview
@Composable
fun TextPagePreview() {
    val navController = rememberNavController()
    TextPage(navController)
}

@Preview
@Composable
fun AppPreview() {
    val navController = rememberNavController()
    NavHost(navController, startDestination = "home") {
        composable("home") { PageScreen(navController) }
        composable("textPage") { TextPage(navController) }
    }
}

@Preview
@Composable
fun AppDarkPreview() {
    val navController = rememberNavController()
    NavHost(navController, startDestination = "home") {
        composable("home") { PageScreen(navController) }
        composable("textPage") { TextPage(navController) }
    }
}

@Composable
fun TextPage(navController: NavHostController) {
    Column (
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp),
        verticalArrangement = Arrangement.Top,
    ) {
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.End
        ) {
            Button(
                onClick = { navController.popBackStack() },
                colors = ButtonDefaults.buttonColors(
                    containerColor = Color.Transparent,
                    contentColor = Color.Gray
                ),
                shape = RectangleShape,
            ) {
//                Text("X", fontsize=10.sp)
                Text("close", fontSize = 15.sp)
            }

        }
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(16.dp),
//        verticalArrangement = Arrangement.SpaceEvenly,
            horizontalAlignment = Alignment.Start
        ) {
            Text("Header1", fontSize = 25.sp, fontWeight = FontWeight.Bold)
            Text("text text text text text text text text text text text text text text" +
                    "text text text text text text text text text text text text text text" +
                    "text text text text text text text text text text text text text text" +
                    "text text text text text text text text text text text text text text")
            Spacer(modifier = Modifier.safeContentPadding())
            Text("Header2", fontSize = 25.sp, fontWeight = FontWeight.Bold)
            Text("text text text text text text text text text text text text text text" +
                    "text text text text text text text text text text text text text text" +
                    "text text text text text text text text text text text text text text" +
                    "text text text text text text text text text text text text text text")
        }
    }
}
