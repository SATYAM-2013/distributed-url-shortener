import streamlit as st
import requests
import os

# ==============================
# Configuration
# ==============================

BACKEND_URL = os.getenv("BACKEND_URL", "http://localhost:8080")
API_KEY = os.getenv("API_KEY", "my-secret-key")  # change this in production

# ==============================
# Page Setup
# ==============================

st.set_page_config(
    page_title="Distributed URL Shortener",
    page_icon="🔗",
    layout="centered"
)

# ==============================
# Custom Styling
# ==============================

st.markdown("""
<style>
.big-title {
    font-size: 40px !important;
    font-weight: 700;
}
.subtitle {
    font-size: 18px;
    color: #9ca3af;
}
.stButton>button {
    border-radius: 10px;
    height: 45px;
    width: 170px;
    font-size: 16px;
}
.result-box {
    background-color: #111827;
    padding: 15px;
    border-radius: 10px;
    word-break: break-all;
}
</style>
""", unsafe_allow_html=True)

# ==============================
# UI
# ==============================

st.markdown('<div class="big-title">🔗 Distributed URL Shortener</div>', unsafe_allow_html=True)
st.markdown('<div class="subtitle">Paste your long URL and get a short link instantly</div>', unsafe_allow_html=True)

st.write("")

long_url = st.text_area(
    "Paste your long URL here:",
    placeholder="Example: https://www.google.com/search?q=distributed+systems"
)

st.write("")

# ==============================
# Shorten Logic
# ==============================

if st.button("Shorten URL"):

    if not long_url.strip():
        st.warning("⚠️ Please enter a valid URL")
    else:
        try:
            response = requests.post(
                f"{BACKEND_URL}/shorten",
                json={"url": long_url},
                headers={
                    "Content-Type": "application/json",
                    "X-API-Key": API_KEY
                },
                timeout=5
            )

            # ------------------------
            # Success
            # ------------------------
            if response.status_code == 200:
                data = response.json()
                short_link = data.get("short_url")

                if short_link:
                    st.success("✅ Short URL Created Successfully")

                    st.markdown(f"""
                    <div class="result-box">
                        <b>Your Short URL:</b><br>
                        {short_link}
                    </div>
                    """, unsafe_allow_html=True)

                    st.code(short_link)
                else:
                    st.error("Invalid response from backend.")

            # ------------------------
            # Unauthorized
            # ------------------------
            elif response.status_code == 401:
                st.error("🔒 Unauthorized - Invalid or missing API key.")

            # ------------------------
            # Rate limit
            # ------------------------
            elif response.status_code == 429:
                st.error("🚦 Rate limit exceeded. Please try again later.")

            # ------------------------
            # Server error
            # ------------------------
            elif response.status_code >= 500:
                st.error("🚨 Server error. Try again later.")

            else:
                st.error(f"❌ Error: {response.status_code}")
                st.text(response.text)

        except requests.exceptions.ConnectionError:
            st.error("🚨 Backend is not running.")

        except requests.exceptions.Timeout:
            st.error("⏳ Request timed out.")

        except Exception as e:
            st.error("Unexpected error occurred.")
            st.text(str(e))

st.write("---")
st.caption("⚠️ This tool generates shortened URLs. Always verify links before sharing.")